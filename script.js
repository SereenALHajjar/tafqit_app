
function converter(s) {
    if (s == "false")
        return false
    return true
}
function sendRequest() {
    const num = document.getElementById('numberInput').value;

    // Get the checked value for feminine/masculine
    const type = document.querySelector('input[name="type"]:checked');
    const feminine = type ? converter(type.value) : false;

    // Get the checked value for hundred (miah/maah)
    const hundred = document.querySelector('input[name="hundred"]:checked');
    const miah = hundred ? converter(hundred.value) : false;

    // Get the checked value for billion (billion/miliar)
    const billion = document.querySelector('input[name="billion"]:checked');
    const billions = billion ? converter(billion.value) : false;

    // Get the checked value for status (nom/ag)
    const status = document.querySelector('input[name="status"]:checked');
    const ag = status ? converter(status.value) : false;

    // Example of how to log or use these values
    console.log({
        num,
        feminine,
        miah,
        billions,
        ag
    });

    // Add your logic to send the request with these values


    if (num === "") {
        alert("Please enter a number");
        return;
    }
    // Create the JSON body for the request
    const requestBody = {
        Num: parseInt(num),
        Feminine: feminine,
        Miah: miah,
        Billions: billions,
        AG: ag
    };

    fetch('http://localhost:8099/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody)
    })
        .then(response => response.text())
        .then(data => {
            document.getElementById('resultOutput').innerText = `${data}`;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}


