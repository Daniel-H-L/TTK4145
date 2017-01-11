//
//  main.c
//  somecode_c
//
//  Created by Mia Olea Vettestad on 11.01.2017.
//  Copyright © 2017 Mia Olea Vettestad. All rights reserved.
//

#include <stdio.h>
#include <pthread.h>

int i = 0;

void* thread1Func() {
    for (int j = 0; j < 1000000; j++) {
        i++;
    }
    return NULL;
}

void* thread2Func() {
    for (int k = 0; k < 1000000; k++) {
        i--;
    }
    return NULL;
}

int main() {
    pthread_t thread1;
    pthread_t thread2;
    
    pthread_create(&thread1, NULL, thread1Func, NULL);
    pthread_create(&thread2, NULL, thread2Func, NULL);
    
    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);
    
    
    printf("%d\n", i);
    
    return 0;
}

