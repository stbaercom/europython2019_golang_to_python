//
// Created by Stefan Baerisch on 2019-07-07.
//

#include "theadsafe_queue.h"

#include <string>
#include <iostream>
#include <thread>

#include <chrono>


using namespace std;


bool goon = true;

struct Task {
    long data1;
    long data2;
};

void push(threadsafe_queue<Task> * queue) {
    long i = 0;
    while(goon) {
        queue->push({i,i});
        i++;
    }
    cout << "Ending Push\n";
}

void pull(threadsafe_queue<Task> * queue) {
    Task t;
    while(goon) {
        queue->wait_and_pop(t);
        if(t.data1 % 10000 == 0) {
            cout << "Got" << t.data1 << "\n";
        }
    }
    cout << "Ending Pull\n";
}




int main()
{
    threadsafe_queue<Task> queue;
    thread t1(push,&queue);
    thread t2(pull,&queue);


    std::this_thread::sleep_for(std::chrono::milliseconds(100));
    goon = false;

    t1.join();
    t2.join();
}