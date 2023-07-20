#include <stdio.h>
#include <stdlib.h>
#include <sys/file.h>
#include <unistd.h>

/* cat but using unix functions rather than std lib */
int main(int argc, char **argv) {
  int fd; // short for file descriptor, which will be a low int identifier of
          // the open file
  int n_bytes_read, n_bites_written;
  char buf[BUFSIZ];

  fd = open(argv[1], O_RDONLY,
            0); // open the file based on what's passed into arg 1
                // read the file as large as BUFSIZ will allow
  while ((n_bytes_read = read(fd, buf, BUFSIZ)) > 0) {
    printf("BUFSIZ is %d\n", BUFSIZ);
    printf("read %d bytes\n", n_bytes_read);
    n_bites_written = write(1, buf, n_bytes_read);
    printf("n_bites_written: %d\n", n_bites_written);
  }
  exit(0);
}
