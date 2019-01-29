export function data() {
    var users;
    fetch(`http://localhost:3100/users`)
    .then(result=>result.json())
    .then(items=>users)
    console.log("hello");
    return users;
}

