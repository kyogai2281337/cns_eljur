export async function makeFetch(url, data = null) {
    try {
        if (data === null) {
            const res = await fetch(url);
            const model = await res.json();
            console.log(model);
            return model
        } else {
            const res = await fetch(url, {
                method: "post",
                credentials: "include",
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });
            const model = await res.json();
            console.log(model);
            return model
        }
        
        
    } catch (err) {
        console.log("Error: ", err);
    }
}


  