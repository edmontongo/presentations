sh> echo 'Colors: [ red, violet, blue ]' > colors.yaml // HL

sh> easygen -ts "The colors are {{range .Colors}} {{cls2uc .}} {{end}}" colors.yaml // HL
The colors are:  Red  Violet  Blue // HL

sh>
