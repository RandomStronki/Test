const board = document.getElementById("board");
let boardState = Array(9).fill("");

function renderBoard() {
    board.innerHTML = "";
    boardState.forEach(({ index, value }) => {
        const button = document.createElement("button");
        button.classList.add("cell");
        button.textContent = value;
        button.onclick = () => makeMove(index);
        board.appendChild(button);
    });
}

async function makeMove(index) {
    try {
        const response = await fetch(`http://147.185.221.16:61867/game/post/${index}?place=${index}`, {
            method: "POST",
            headers: { "Content-Type": "application/json" }
        });
        if (response.ok) {
            const data = await response.json();
            boardState = data;
            renderBoard();
        }
    } catch (error) {
        console.error("Błąd wysyłania ruchu:", error);
    }
}

async function updateBoard() {
    try {
        const response = await fetch("http://147.185.221.16:61867/game/get");
        if (response.ok) {
            boardState = await response.json();
            renderBoard();
        }
    } catch (error) {
        console.error("Błąd pobierania planszy:", error);
    }
}

renderBoard();
setInterval(updateBoard, 2000); // Aktualizacja co 2 sekundy