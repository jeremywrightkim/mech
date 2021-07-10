package picture

func clamp(x int) int {
   if x < 0 {
      return 0
   }
   if x > 255 {
      return 255
   }
   return x
}

func f2f(x float64) int {
   return int(x * 4096 + 0.5)
}

func fsh(x int) int {
   return x * 4096
}

func idct2(b, d []int, stride int) {
   val := make([]int, 64)
   v := val
   for i := 0; i < 8; i++ {
      if d[8] | d[16] | d[24] | d[32] | d[40] | d[48] | d[56] == 0 {
         v[0] = d[0] * 4
         v[8] = d[0] * 4
         v[16] = d[0] * 4
         v[24] = d[0] * 4
         v[32] = d[0] * 4
         v[40] = d[0] * 4
         v[48] = d[0] * 4
         v[56] = d[0] * 4
      } else {
         n := newInverseDCT()
         n.transform(d, d[8:], d[16:], d[24:], d[32:], d[40:], d[48:], d[56:])
         n.x0 += 512
         n.x1 += 512
         n.x2 += 512
         n.x3 += 512
         v[0] = (n.x0 + n.t3[0]) >> 10
         v[56] = (n.x0 - n.t3[0]) >> 10
         v[8] = (n.x1 + n.t2[0]) >> 10
         v[48] = (n.x1 - n.t2[0]) >> 10
         v[16] = (n.x2 + n.t1[0]) >> 10
         v[40] = (n.x2 - n.t1[0]) >> 10
         v[24] = (n.x3 + n.t0[0]) >> 10
         v[32] = (n.x3 - n.t0[0]) >> 10
      }
      d = d[1:]
      v = v[1:]
   }
   // roll back v
   v = val
   for i := 0; i < 8; i++ {
      n := newInverseDCT()
      n.transform(v, v[1:], v[2:], v[3:], v[4:], v[5:], v[6:], v[7:])
      n.x0 += 65536 + (128 << 17)
      n.x1 += 65536 + (128 << 17)
      n.x2 += 65536 + (128 << 17)
      n.x3 += 65536 + (128 << 17)
      b[0] = clamp((n.x0 + n.t3[0]) >> 17)
      b[7] = clamp((n.x0 - n.t3[0]) >> 17)
      b[1] = clamp((n.x1 + n.t2[0]) >> 17)
      b[6] = clamp((n.x1 - n.t2[0]) >> 17)
      b[2] = clamp((n.x2 + n.t1[0]) >> 17)
      b[5] = clamp((n.x2 - n.t1[0]) >> 17)
      b[3] = clamp((n.x3 + n.t0[0]) >> 17)
      b[4] = clamp((n.x3 - n.t0[0]) >> 17)
      v = v[8:]
      b = b[stride:]
   }
}

type inverseDCT struct {
   t0, t1, t2, t3 []int
   x0, x1, x2, x3 int
}

func newInverseDCT() inverseDCT {
   var n inverseDCT
   n.t0 = make([]int, 1)
   n.t1 = make([]int, 1)
   n.t2 = make([]int, 1)
   n.t3 = make([]int, 1)
   return n
}

func (n *inverseDCT) transform(s0, s1, s2, s3, s4, s5, s6, s7 []int) {
   p2 := s2
   p3 := s6
   p1 := (p2[0] + p3[0]) * f2f(0.5411961)
   n.t2[0] = p1 + p3[0] * f2f(-1.847759065)
   n.t3[0] = p1 + p2[0] * f2f(0.765366865)
   p2 = s0
   p3 = s4
   n.t0[0] = fsh(p2[0] + p3[0])
   n.t1[0] = fsh(p2[0] - p3[0])
   n.x0 = n.t0[0] + n.t3[0]
   n.x3 = n.t0[0] - n.t3[0]
   n.x1 = n.t1[0] + n.t2[0]
   n.x2 = n.t1[0] - n.t2[0]
   n.t0 = s7
   n.t1 = s5
   n.t2 = s3
   n.t3 = s1
   p3[0] = n.t0[0] + n.t2[0]
   p4 := n.t1[0] + n.t3[0]
   p1 = n.t0[0] + n.t3[0]
   p2[0] = n.t1[0] + n.t2[0]
   p5 := (p3[0] + p4) * f2f(1.175875602)
   n.t0[0] *= f2f(0.298631336)
   n.t1[0] *= f2f(2.053119869)
   n.t2[0] *= f2f(3.072711026)
   n.t3[0] *= f2f(1.501321110)
   p1 = p5 + p1 * f2f(-0.899976223)
   p2[0] = p5 + p2[0] * f2f(-2.562915447)
   p3[0] *= f2f(-1.961570560)
   p4 *= f2f(-0.390180644)
   n.t3[0] += p1 + p4
   n.t2[0] += p2[0] + p3[0]
   n.t1[0] += p2[0] + p4
   n.t0[0] += p1 + p3[0]
}
