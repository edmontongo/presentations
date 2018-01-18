#include <stdio.h>
#include <errno.h>

int main() {
	FILE *f = fopen("/etc/passwd", "r");
	if (f == NULL) {
		// handle error
		fprintf(stderr, "File opening failed: %d\n", errno);
		return 1;
	}
	// ...use the file...

	fclose(f);
	return 0;
}
