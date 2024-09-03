FROM mcr.microsoft.com/oss/go/microsoft/golang:1.23-fips-cbl-mariner2.0 AS builder
WORKDIR /source

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o /app/sum . ; exit 1 # The command did not fail, so exit 1 to keep the log around.

# > [6/6] RUN go build -o /app/sum . ; exit 1 # The command did not fail, so exit 1 to keep the log around.:
# 10.05 # vendor/github.com/golang-fips/openssl/v2
# 10.05 thread_setup_unix.c: In function 'thread_id':
# 10.05 thread_setup_unix.c:32:12: warning: 'pthread_setspecific' expecting 1 byte in a region of size 0 [-Wstringop-overread]
# 10.05    32 |     (void) pthread_setspecific(destructor_key, (void*)1);
# 10.05       |            ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# 10.05 In file included from thread_setup_unix.c:5:
# 10.05 /usr/lib/gcc/x86_64-pc-linux-gnu/11.2.0/include-fixed/pthread.h:1317:12: note: in a call to function 'pthread_setspecific' declared with attribute 'access (none, 2)'
# 10.05  1317 | extern int pthread_setspecific (pthread_key_t __key,
# 10.05       |            ^~~~~~~~~~~~~~~~~~~
