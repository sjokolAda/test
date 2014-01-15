// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread


#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lock;

// Note the return type: void*
void* adder(){
    pthread_mutex_lock(&lock);
    for(int x = 0; x < 1000000; x++){
        i++;
    }
    pthread_mutex_unlock(&lock);
    return NULL;
}

void* subtractor(){
    pthread_mutex_lock(&lock);
    for(int x = 0; x < 1000000; x++){
        i--;
    }
    pthread_mutex_unlock(&lock);
    return NULL;
}


int main(){

    if (pthread_mutex_init(&lock, NULL) != 0)
    {
        printf("\n mutex init failed\n");
        return 1;
    }


    pthread_t adder_thr;
    pthread_t subtractor_thr;
    pthread_create(&adder_thr, NULL, adder, NULL);
    

    pthread_create(&subtractor_thr, NULL, subtractor, NULL);
    
    for(int x = 0; x < 50; x++){
        printf("%i\n", i);
    }

    
    pthread_join(adder_thr, NULL);
    pthread_join(subtractor_thr, NULL);
    printf("Done: %i\n", i);
    pthread_mutex_destroy(&lock);
    return 0;
    
}
