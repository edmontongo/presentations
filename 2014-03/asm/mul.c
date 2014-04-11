#include "runtime.h"

void ·Mul(int64 a, int64 b, int64 res) {
    res = a * b;
    FLUSH(&res);
}
