package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// เชื่อมต่อกับเซิร์ฟเวอร์
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	// ปิดการเชื่อมต่อเมื่อแอปพลิเคชันจบการทำงาน
	defer conn.Close()

	fmt.Println("เชื่อมต่อกับเซิร์ฟเวอร์แล้ว")

	reader := bufio.NewReader(os.Stdin)
	for {
		// อ่านข้อมูลจากผู้ใช้
		fmt.Print("ป้อนข้อความ: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// ส่งข้อความไปยังเซิร์ฟเวอร์
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// พิมพ์จำนวนบายต์ที่ส่ง
		fmt.Printf("ส่ง %d บายต์\n", len(message))

		// รับและพิมพ์การตอบกลับจากเซิร์ฟเวอร์
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("การตอบกลับจากเซิร์ฟเวอร์: %s", buffer[:n])
	}
}
