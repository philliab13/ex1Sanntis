// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>


int i = 0;

pthread_mutex_t lock; 

// Note the return type: void*
void* incrementingThreadFunction(){
    
    for (int j = 0; j < 1000000; j++)
    {
        pthread_mutex_lock(&lock);
        i++;
        pthread_mutex_unlock(&lock); 
    }
    
    return NULL;
}

void* decrementingThreadFunction(){
    
    for (int j = 999999; j > 0; j--)
    {
        pthread_mutex_lock(&lock);
        i--;
        pthread_mutex_unlock(&lock); 
    }
    
    
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?

    pthread_t ptid;
    pthread_t ptid2;

    if (pthread_mutex_init(&lock, NULL) != 0) { 
        printf("\n mutex init has failed\n"); 
        return 1;   
    } 

    pthread_create(&ptid, NULL, &incrementingThreadFunction, NULL); 
    pthread_create(&ptid2, NULL, &decrementingThreadFunction, NULL); 
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    

    pthread_join(ptid, NULL); 
    pthread_join(ptid2, NULL);
    pthread_mutex_destroy(&lock);
    
    printf("The magic number is: %d\n", i);
    return 0;
}
