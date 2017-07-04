# golang if, switch, map performance benchmark

## Device spec
### Test1 
  MODEL: Apple Macbook Pro Retina 2014 mid CPU CUSTOM  
  CPU: Intel i7-4980HQ  
  RAM: 16GB  
  OS: Mac OSX 10.12.5 64bit  
### Test2
  MODEL: VMWare ESXI 6.5a Virtualized   
  CPU: Intel Xeon E3 1245v5 2 Virtualized CPU   
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
    BenchmarkSwitchIntShort-2    	100000000	        12.9 ns/op
    BenchmarkIfIntShort-2        	100000000	        10.5 ns/op
    BenchmarkMapIntShort-2       	100000000	        17.2 ns/op
    BenchmarkSwitchFuncShort-2   	100000000	        12.8 ns/op
    BenchmarkIfFuncShort-2       	100000000	        10.6 ns/op
    BenchmarkMapFuncShort-2      	100000000	        17.2 ns/op
    BenchmarkSwitchIntLong-2     	50000000	        26.8 ns/op
    BenchmarkIfIntLong-2         	100000000	        15.0 ns/op
    BenchmarkMapIntLong-2        	100000000	        19.2 ns/op
    BenchmarkSwitchFuncLong-2    	50000000	        26.8 ns/op
    BenchmarkIfFuncLong-2        	100000000	        15.1 ns/op
    BenchmarkMapFuncLong-2       	100000000	        19.9 ns/op
    PASS
    ok  	command-line-arguments	17.943s

## Conclusion
  Why virtual machine is faster then my real mac?
