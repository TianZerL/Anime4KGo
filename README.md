# Anime4KGo
This is an implementation of Anime4K in Go. It based on the [bloc97's Anime4K](https://github.com/bloc97/Anime4K) algorithm version 0.9 and some optimizations have been made.  
This project is for learning and the exploration task of algorithm course in SWJTU.

# About Anime4K
Anime4K is a simple high-quality anime upscale algorithm for anime. it does not use any machine learning approaches, and can be very fast in real-time processing.

# Usage
    -?    Show help information
    -f    Fast Mode but low quality
    -h    Show help information
    -i string
            File for loading (default "./pic/p1.png")
    -o string
            File for outputting (default "out.png")
    -p int
            Passes for processing (default 2)
    -sc float
            Strength for pushing color,range 0 to 1,higher for thinner (default 0.3333333333333333)
    -sg float
            Strength for pushing gradient,range 0 to 1,higher for sharper (default 1)

# Other implementations
- Python
  - [TianZerL/Anime4KPython](https://github.com/TianZerL/Anime4KPython)
- C++
  - [TianZerL/Anime4KCPP](https://github.com/TianZerL/Anime4KCPP)
- C#
  - [shadow578/Anime4kSharp](https://github.com/shadow578/Anime4kSharp)
  - [net2cn/Anime4KSharp](https://github.com/net2cn/Anime4KSharp)
- Java
  - [bloc97/Anime4K](https://github.com/bloc97/Anime4K)
- Rust
  - [andraantariksa/Anime4K-rs](https://github.com/andraantariksa/Anime4K-rs)