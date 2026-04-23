package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"password_manager/crypto"
	"password_manager/storage"
	"password_manager/vault"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: add | list")
		return
	}

	fmt.Print("Enter master password: ")
	bytePw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	fmt.Println()

	salt := []byte("fix-salt-32-bytes-length-at-least")
	key := crypto.DeriveKey(string(bytePw), salt)

	v := vault.NewVault()
	data, err := storage.Load()
	if err == nil {
		decrypted, err := crypto.Decrypt(data, key)
		if err == nil {
			v.FromBytes(decrypted)
		}
	}

	switch os.Args[1] {
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Resource: ")
		res, _ := reader.ReadString('\n')
		fmt.Print("Login: ")
		log, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		pass, _ := reader.ReadString('\n')

		v.AddAccount(vault.Account{
			Resource: strings.TrimSpace(res),
			Login:    strings.TrimSpace(log),
			Password: strings.TrimSpace(pass),
		})

		payload, _ := v.ToBytes()
		encrypted, _ := crypto.Encrypt(payload, key)
		storage.Save(encrypted)
		fmt.Println("Saved!")

	case "list":
		for _, acc := range v.Accounts {
			fmt.Printf("Res: %s | Login: %s | Pass: %s\n", acc.Resource, acc.Login, acc.Password)
		}
	}
}
