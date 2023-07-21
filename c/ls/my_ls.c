// Please implement a minimal clone of the ls program. We have chosen this
// exercise as it will require you to use structs, pointers and arrays, as well
// as some C standard library functions with interesting interfaces. It will
// also likely to be substantial enough to merit some degree of code
// organization. Minimally, it should list the contents of a directory including
// some information about each file, such as file size. As a stretch goal, use
// man ls to identify any interesting flags you may wish to support, and
// implement them.
#include <dirent.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <libgen.h>

#define MAX_PATH 1024

void printdir(char *);
void dirwalk(char *, void (*)(char *));

int main(int argc, char **argv) {

  if (argc == 1)
    printdir(".");
  else
    while (--argc > 0)
      printdir(*++argv);

  return 0;
}

void printdir(char *name) {
  struct stat stbuf;

  if (stat(name, &stbuf) == -1) {
    fprintf(stderr, "printdir: can't access %s\n", name);
    return;
  }
  if ((stbuf.st_mode & S_IFMT) == S_IFDIR)
    // dirwalk(name, printdir);
    printf("%s/", name);
  if (basename(name)[0] != '.')
    printf("%s\n", basename(name));
}

void dirwalk(char *dir, void (*fcn)(char *)) {
  char name[MAX_PATH];
  struct dirent *dp;
  DIR *dfd;

  if ((dfd = opendir(dir)) == NULL) {
    fprintf(stderr, "dirwalk: can't open %s\n", dir);
    return;
  }

  while ((dp = readdir(dfd)) != NULL) {
    if (strcmp(dp->d_name, ".") == 0 || strcmp(dp->d_name, "..") == 0)
      continue;
    if (strlen(dir) + strlen(dp->d_name) + 2 > sizeof(name))
      fprintf(stderr, "dirwalk: name %s/%s too long\n", dir, dp->d_name);
    else {
      sprintf(name, "%s/%s", dir, dp->d_name);
      (*fcn)(name);
    }
  }
  closedir(dfd);
}
