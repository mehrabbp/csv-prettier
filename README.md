## What's a csv pretty printer?
In short, a csv pretty printer will read a csv file as input and print it like a table for output:
### Input
```
post_id,phone,title,description,device_id
1,09122965123,oven,clean and perfect,3d345kfx
2,09382461163,fridge,beautiful & spacious & assured,xcvd932f
3,09322935123,desk,just a wooden desk,i3xd3z3o
```

### Output
```
+-------------------------------------------------------------------+
|post_id|phone      |title |description                   |device_id|
+-------------------------------------------------------------------+
|1      |09122965123|oven  |clean and perfect             |3d345kfx |
+-------------------------------------------------------------------+
|2      |09382461163|fridge|beautiful & spacious & assured|xcvd932f |
+-------------------------------------------------------------------+
|3      |09322935123|desk  |just a wooden desk            |i3xd3z3o |
+-------------------------------------------------------------------+
```

# Part I: Prologue
Implement a csv pretty printer for our clients! Your program should take one argument as input which is the csv file path and output to stdout the pretty printed version. For example:
`./csv-pprint.out example.csv`
### Constraints

* Each line has a fixed number of commas.
* Each line is less than 80 characters.


============================================================
============================================================

Part 2:


Our clients have reported that they need to use our program for files with longer lines. 
But they don't want to scroll horizontally in order to see the full content of a line.
Therefore, they want to specify the max width of `80` for each column. Here is an example:

### Input
```
post_id,phone,title,description,device_id
1,09122965123,oven,clean and perfect oven for baking cakes & making pizza & boiling canned food. Warning: ovens are not for programmers,3d345kfx
2,09382461163,fridge,beautiful & spacious & assured this is the Emersan fridge. for more information please call 911,xcvd932f
3,09322935123,desk,just a wooden desk,i3xd3z3o
```

### Output
```
+---------------------------------------------------------------------------------------------------------------------+
|post_id|phone      |title |description                                                                     |device_id|
+---------------------------------------------------------------------------------------------------------------------+
|1      |09122965123|oven  |clean and perfect oven for baking cakes & making pizza & boiling canned food.   |3d345kfx |
|       |           |      |Warning: ovens are not for programmer                                           |         |
+---------------------------------------------------------------------------------------------------------------------+
|2      |09382461163|fridge|beautiful & spacious & assured this is the Emersan fridge. for more information |xcvd932f |
|       |           |      |please call 911                                                                 |         |
+---------------------------------------------------------------------------------------------------------------------+
|3      |09322935123|desk  |just a wooden desk                                                              |i3xd3z3o |
+---------------------------------------------------------------------------------------------------------------------+
```

--------

```
post_id,phone,title,description,device_id
1,09122965123,oven,clean and perfect oven for baking cakes & making pizza & boiling canned food. Warning: ovens are not for programmers,3d345kfx
2,09382461163,fridge,beautiful & spacious & assured this is the Emersan fridge. for more information please call 911,xcvd932f
3,09322935123,desk,just a wooden desk,i3xd3z3o
4,01234567890,num,12345678901234567890123456789012345678901 12345678901234567890123456789012345678901 12345678901234567890123456789012345678901,3z3oi3xd
```

### Output
```
+---------------------------------------------------------------------------------------------------------------------+
|post_id|phone      |title |description                                                                     |device_id|
+---------------------------------------------------------------------------------------------------------------------+
|1      |09122965123|oven  |clean and perfect oven for baking cakes & making pizza & boiling canned food.   |3d345kfx |
|       |           |      |Warning: ovens are not for programmer                                           |         |
+---------------------------------------------------------------------------------------------------------------------+
|2      |09382461163|fridge|beautiful & spacious & assured this is the Emersan fridge. for more information |xcvd932f |
|       |           |      |please call 911                                                                 |         |
+---------------------------------------------------------------------------------------------------------------------+
|3      |09322935123|desk  |just a wooden desk                                                              |i3xd3z3o |
+---------------------------------------------------------------------------------------------------------------------+
|4      |01234567890|num   |12345678901234567890123456789012345678901                                       |3z3oi3xd |
|       |           |      |12345678901234567890123456789012345678901                                       |         |
|       |           |      |12345678901234567890123456789012345678901                                       |         |
+---------------------------------------------------------------------------------------------------------------------+
```



============================================================
============================================================


Part 3:

Our clients are making us MAD! They want us to make vertical alignment!
For example if our cell output is one line but the cell height is 3 lines,
we should print the line in the center of the cell, not the top of it. Please see the example for more details.

Input

```
post_id,phone,title,description,device_id
1,09122965123,oven,clean and perfect oven for baking cakes - making pizza - boiling canned food. Warning: ovens are not for programmers blah blah blah blah blah blah blah blah blah,3d345kfx
2,09382461163,fridge,beautiful & spacious & assured this is the Emersan fridge. for more information please call 911,xcvd932f
3,09322935123,desk,just a wooden desk,i3xd3z3o
```


Output

```
+---------------------------------------------------------------------------------------------------------------------+
|post_id|phone      |title |description                                                                     |device_id|
+---------------------------------------------------------------------------------------------------------------------+
|       |           |      |clean and perfect oven for baking cakes - making pizza - boiling canned food.   |         |
|1      |09122965123|oven  |Warning: ovens are not for programmer blah blah blah blah blah blah blah blah   |3d345kfx |
|       |           |      |blah                                                                            |         |
+---------------------------------------------------------------------------------------------------------------------+
|2      |09382461163|fridge|beautiful & spacious & assured this is the Emersan fridge. for more information |xcvd932f |
|       |           |      |please call 911                                                                 |         |
+---------------------------------------------------------------------------------------------------------------------+
|3      |09322935123|desk  |just a wooden desk                                                              |i3xd3z3o |
+---------------------------------------------------------------------------------------------------------------------+
```



============================================================
============================================================

Part 4:


Now our clients are starting to use our tool instead of excel! Now they want to sort the table before printing based on some column. Please add such functionality using execution arguments. For example:
./csv-pprint.out --sort phone --max-width 40 file.csv

Input
```
post_id,phone,title,description,device_id
1,09122965123,oven,clean and perfect oven for baking cakes - making pizza - boiling canned food. Warning: ovens are not for programmers blah blah blah blah blah blah blah blah blah,3d345kfx
2,09382461163,fridge,beautiful & spacious & assured this is the Emersan fridge. for more information please call 911,xcvd932f
3,09322935123,desk,just a wooden desk,i3xd3z3o
```


Output
```
The output is now sorted based on phone.
+---------------------------------------------------------------------------------------------------------------------+
|post_id|phone      |title |description                                                                     |device_id|
+---------------------------------------------------------------------------------------------------------------------+
|       |           |      |clean and perfect oven for baking cakes - making pizza - boiling canned food.   |         |
|1      |09122965123|oven  |Warning: ovens are not for programmer blah blah blah blah blah blah blah blah   |3d345kfx |
|       |           |      |blah                                                                            |         |
+---------------------------------------------------------------------------------------------------------------------+
|3      |09322935123|desk  |just a wooden desk                                                              |i3xd3z3o |
+---------------------------------------------------------------------------------------------------------------------+
|2      |09382461163|fridge|beautiful & spacious & assured this is the Emersan fridge. for more information |xcvd932f |
|       |           |      |please call 911                                                                 |         |
+---------------------------------------------------------------------------------------------------------------------+
```



============================================================
============================================================

part 5:
 
 
Our clients have reported that some of their data are raw text content which include the comma character. Therefore they will use double quotes for fields that have comma in their values. Your new version of program should handle such files.

Input
```
post_id,phone,title,description,device_id
1,09122965123,oven,"clean and perfect oven for baking cakes , making pizza , boiling canned food. Warning: ovens are not for programmers",3d345kfx
2,09382461163,fridge,beautiful & spacious & assured this is the Emersan fridge. for more information please call 911,xcvd932f
3,09322935123,desk,just a wooden desk,i3xd3z3o
```


Output
```
+---------------------------------------------------------------------------------------------------------------------+
|post_id|phone      |title |description                                                                     |device_id|
+---------------------------------------------------------------------------------------------------------------------+
|1      |09122965123|oven  |clean and perfect oven for baking cakes , making pizza , boiling canned food.   |3d345kfx |
|       |           |      |Warning: ovens are not for programmer                                           |         |
+---------------------------------------------------------------------------------------------------------------------+
|2      |09382461163|fridge|beautiful & spacious & assured this is the Emersan fridge. for more information |xcvd932f |
|       |           |      |please call 911                                                                 |         |
+---------------------------------------------------------------------------------------------------------------------+
|3      |09322935123|desk  |just a wooden desk                                                              |i3xd3z3o |
+---------------------------------------------------------------------------------------------------------------------+
```
