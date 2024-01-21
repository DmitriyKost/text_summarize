async function onSubmit() {
    let contents = document.getElementById("textArea").value;
    sendText(contents).then(
        yaay => { 
            let outField = document.getElementById("outText");
            outField.textContent="Чет";
            outField.style.visibility = "visible";
            console.log(yaay);
        },
        fuck => {
            document.getElementById("Success").style.visibility = "visible";

            console.log(fuck);
        }
    );
}

async function sendText(text) {
    let res = fetch("/summarize", {
        method: "POST",
        body: text,
    });
    return res
}
