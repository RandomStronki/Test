const board = document.getElementById("board");
let boardState = Array.from({ length: 9 }, (_, i) => ({ index: i, value: "" }));

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
        console.log(`Wysyłanie POST na /game/post/${index}`);
        const response = await fetch(`http://127.0.0.1:8080/game/post/${index}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify({ index })
        });

        console.log("Response status:", response.status);
        if (!response.ok) {
            throw new Error(`Serwer zwrócił błąd: ${response.status}`);
        }

        const data = await response.json();
        boardState = data;
        renderBoard();
    } catch (error) {
        console.error("Błąd wysyłania ruchu:", error);
    }
}

async function updateBoard() {
    try {
        const response = await fetch("http://127.0.0.1:8080/game/get");
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
