#include "gohttplib.h"

void handler(void* w)
{
    char* buf = "Hello, world";
    ResponseWriter_Write(w, buf, 12);
    // TODO: Handle return
}

int main()
{
    HandleFunc("/hello", &handler);
    ListenAndServe(":8000");
    return 0;
}
