package main

func reverse(x int) int {
    res := 0
    sign := 1
    if x < 0 {
        sign = -1
        x *= -1
    }
    
    for {
        res = res * 10 + x % 10
        x /= 10
        if x <= 0 {
            break
        }
    }
    if res >= int(int32(^uint32(0) >> 1)) { //where ^uint32(0) is a reversed uint32 = 111111..111 
                                //shifted right one bit which gives maximum int32
        return 0
    }
    return res * sign
}