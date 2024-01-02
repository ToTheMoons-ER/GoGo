// แพ็คเกจหลักของโปรแกรม ตั้งชื่อว่า main
package main

// นำเข้าแพ็คเกจที่ต้องการใช้งาน
import (
	"fmt"
	"net"
)

// ฟังก์ชัน handleConnection ใช้ในการจัดการการเชื่อมต่อ
func handleConnection(conn net.Conn) {
	defer conn.Close() // ปิดการเชื่อมต่อก่อนที่จะออกจากฟังก์ชัน

	// buffer สำหรับการอ่าน
	buffer := make([]byte, 1024)
	for {
		// อ่านข้อมูลจาก client
		n, err := conn.Read(buffer) // Read() จะบล็อกจนกว่าจะได้ข้อมูลจากเครือข่ายและ n คือจำนวนบายต์ที่อ่านได้
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		// พิมพ์จำนวนบายต์ที่อ่านได้
		fmt.Printf("Received %d bytes\n", n)

		// พิมพ์ข้อมูลที่ได้รับ
		// [0 : n], [:n], [0:n] เป็นรูปแบบเท่ากัน
		fmt.Printf("Received message: %s", buffer[:n]) // :n เป็นตัวดัชนีที่ให้สเลิซท์และคืนสเลิซของบายต์แรกถึงบายต์ที่ n

		// พิมพ์ข้อความเป็นบายต์
		fmt.Printf("Received message as bytes: %v\n", buffer[:n])

		// 104 คือรหัส ASCII ของ 'h'
		// 105 คือรหัส ASCII ของ 'i'
		// 10 คือรหัส ASCII ของ '\n'

		// ส่งข้อความยืนยันกลับไปยัง client
		response := "Message received successfully\n"
		conn.Write([]byte(response))
	}
}

func main() {
	// 1. สร้างตัวแปร listener เพื่อรับการเชื่อมต่อเครือข่าย
	// โดยใช้ฟังก์ชัน net.Listen และกำหนดพอร์ตที่ต้องการให้เปิดใช้งาน
	// ในที่นี้ใช้พอร์ต 5000
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	// ปิด Listener เมื่อปิดแอปพลิเคชัน
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	// รอการเชื่อมต่อเข้ามา
	for {
		// Accept() จะบล็อกจนกว่าจะมีการเชื่อมต่อ
		conn, err := listener.Accept() // Accept เป็นกระบวนการทำแฮนด์เชคสามทาง
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue // ไปที่การวนลูปถัดไป
		}

		fmt.Println("New connection established")

		// จัดการการเชื่อมต่อในกอรูทีนใหม่
		go handleConnection(conn)
	}
}
