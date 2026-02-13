#include "fast_logic.h"
#include <stdio.h>

void process_data(int *arr, int size) {
    for (int i = 0; i < size; i++) {
        arr[i] = arr[i] * 2;
    }
}
