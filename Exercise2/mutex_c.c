//
//  mutex_c.c
//  
//
//  Created by Mia Olea Vettestad on 18.01.2017.
//
//

#include "mutex_c.h"
#include <stdio.h>
#include <pthread.h>

int i = 0;

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

void* thread1Func() {
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&mutex);
        i++;
        pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void* thread2Func() {
    for (int k = 0; k < (1000000 - 1); k++) {
        pthread_mutex_lock(&mutex);
        i--;
        pthread_mutex_unlock(&mutex);
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
