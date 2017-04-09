let socket = new WebSocket("ws://localhost:3000/wsserver/kelvin2")

socket.onmessage = (msg) => {
    console.log(msg.data)
}

socket.onerror = (err) => {
    console.log(err)
}

socket.onopen = () => {
    console.log("Conexão estabelecida")
}

socket.onclose = (motive) => {
    console.log(motive)
}
