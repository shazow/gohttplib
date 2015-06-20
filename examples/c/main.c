#include <stdio.h>
#include "gohttplib.h"

char * handler()
{
    return "Hello, world.";
}

int main()
{
    HandleFunc("/hello", &handler);
    ListenAndServe(":8000");
    return 0;
}
