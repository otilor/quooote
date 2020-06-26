document.addEventListener("keypress", function (event) {

    if (event.ctrlKey && event.keyCode === 13) {
        submitForm();

    }
});

function submitForm () {
    let quoteTitle = document.getElementById("quote_title").value;
    let quoteBody = document.getElementById("quote_body").value;
    if (quoteTitle.length > 0 && quoteBody.length > 0) {
        document.quooote_form.submit();
    }
}

function deleteQuote(obj, abc) {
    let quoteBody = document.getElementById("quote-body");
    quoteBody.className += "strikethrough";
    startCountDown();

}

function startCountDown() {
    let counter = 10;
    id = document.getElementById("counter");
    id.innerHTML = counter;

    setInterval( function () {
        counter--;

        if (counter > 0 ){
            id.innerHTML = counter;

        }
        if (counter === 0) {
            window.location.href = "/";
        }
    }, 1000);

}
