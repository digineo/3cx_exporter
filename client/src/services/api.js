
export async function GetStatus(){
    const resp = await fetch("/status",{method:"GET"})
    return await resp.json()
}

export async function GetConfig(){
    const resp = await fetch("/config",{method:"GET"})
    return await resp.json()
    
}

export async function SetConfig(host,username,password){
    const resp = await fetch("/config",{
        method:"PUT",
        body: JSON.stringify({
            Hostname: host,
            Username: username,
            Password: password

        })
    })
    const rr = await resp.json()
    if (resp.status!=200){
        throw `Settings not saved. Message ${rr.error}`

    }
    return await resp.json()
}
