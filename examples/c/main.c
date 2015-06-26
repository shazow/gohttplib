#include <stdio.h>
#include "../../build/libgohttp.h"

void handler(ResponseWriter *w, Request *req)
{
    printf("handler: %s %s\n", req->Method, req->URL);

    char *buf = "Hello, world";
    int n = ResponseWriter_Write(w, buf, 12);

    if (n == EOF) { 
        printf("handler: Failed to write.\n");
        ResponseWriter_WriteHeader(w, 500);
        return;
    }

    printf("handler: Wrote %d bytes.\n", n);
}

int main()
{
    HandleFunc("/hello", &handler);
    printf("Listening on :8000\n");
    ListenAndServe(":8000");
    return 0;
}
