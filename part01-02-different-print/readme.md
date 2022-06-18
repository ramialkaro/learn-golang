## Different Print funcs and format in Go

Using fmt package to handle I/O (Input / Output) from Golang to print the output on console. 

To learn more about the package, here is the [Offical docs](https://pkg.go.dev/fmt)

----

#### Available commands
To see the output or run the file, you could use ``` go run . ``` (to run all files) or could specifie the file name for example ``` go run main.go ```

---
#### Output will be

<div style="background-color: black; padding: 1.3em; color:green;">
Using the normal Print func.
With Println it have the same behavior as Print except this func will add break line at the of the text.
<br>You could specifi any format that you want, for exmaple
<br> decimal 42384
<br>float 42.234000
<br>float with 2 percision 42.23
</div>

---
#### Different print [format](https://pkg.go.dev/fmt)
Different print format for the func. for example decimal and float
```
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
```
---
#### Other [func](https://pkg.go.dev/fmt#pkg-functions) under fmt packge

There are different print func under fmt package. for example 

* Errorf
* Fprint
* Fprintf
* Fprintln
* Fscan
* Fscanf
* Fscanln
* Print
* Printf
* Println
* Scan
* Scanf
* Scanln
* Sprint
* Sprintf
* Sprintln
* Sscan
* Sscanf
* Sscanln
