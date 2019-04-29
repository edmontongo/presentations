#!/bin/bash
rm *png
ls *ditaa | xargs -I {} ditaa -A -T -S -s 1.5 {}
