const client = 'kelvin'

let socket = new WebSocket(`ws://localhost:3000/wsserver/${client}`)

let message = document.getElementById("message")
let clientId = document.getElementById("clientId")

document.getElementById("send").addEventListener('click', () => {

    socket.send(JSON.stringify({
            From: client,
            Data: message.value,
            Destiny: clientId.value
    }))

    message.value = ""
})

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
