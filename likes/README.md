# Realtime Clap Counter

Just a simple implementation to show realtime clap counter. Made using gorilla websocket.

## Requirements
- Go >= 1.6

## Instructions
- Server can be started using `make setup run`
- Go to [localhost:8080/](http://localhost:8080/) and click on clap image to see it in action. Try on multiple tabs or browsers

## Performance
- Tested on Ubuntu Box with following config, and initiated ~ 50,000 client connections but only 5,000 clients were able to connect
```
Architecture:                    x86_64
CPU op-mode(s):                  32-bit, 64-bit
Byte Order:                      Little Endian
Address sizes:                   39 bits physical, 48 bits virtual
CPU(s):                          2
On-line CPU(s) list:             0,1
Thread(s) per core:              1
Core(s) per socket:              2
Socket(s):                       1
NUMA node(s):                    1
Vendor ID:                       GenuineIntel
CPU family:                      6
Model:                           140
Model name:                      11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
Stepping:                        1
CPU MHz:                         2803.198
BogoMIPS:                        5606.39
Hypervisor vendor:               KVM

```
![image](https://user-images.githubusercontent.com/3920286/132701330-43f8171c-fe87-49f4-8dd0-663bc46270fb.png)
