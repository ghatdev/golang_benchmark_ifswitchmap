# golang if, switch, map performance benchmark

## Device spec
### Test1 
  MODEL: Apple Macbook Pro Retina 2014 mid CPU CUSTOM  
  CPU: Intel i7-4980HQ  
  RAM: 16GB  
  OS: Mac OSX 10.12.5 64bit  
### Test2
  MODEL: VMWare ESXI 6.5a Virtualized 
  CPU: Intel Xeon E3 2 Virtualized CPU 
  RAM: 2GB  
  OS: Ubuntu Linux 64bit 17.04 Desktop  

## Result 
### Test1
    BenchmarkSwitchIntShort-8    	100000000	        14.3 ns/op
    BenchmarkIfIntShort-8        	100000000	        11.5 ns/op
    BenchmarkMapIntShort-8       	100000000	        17.3 ns/op
    BenchmarkSwitchFuncShort-8   	100000000	        14.4 ns/op
    BenchmarkIfFuncShort-8       	100000000	        11.4 ns/op
    BenchmarkMapFuncShort-8      	100000000	        17.8 ns/op
    BenchmarkSwitchIntLong-8     	50000000	        27.6 ns/op
    BenchmarkIfIntLong-8         	100000000	        15.8 ns/op
    BenchmarkMapIntLong-8        	100000000	        23.7 ns/op
    BenchmarkSwitchFuncLong-8    	50000000	        27.3 ns/op
    BenchmarkIfFuncLong-8        	100000000	        15.8 ns/op
    BenchmarkMapFuncLong-8       	50000000	        25.0 ns/op
    PASS
    ok  	command-line-arguments	18.436s
      
### Test2

## Conclusion
  NOT YET
