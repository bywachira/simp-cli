# simp-cli
A simple code execution tools that does whatever you want it to just not keep-alive type of commands

## How it works

Create a json file in the root of your project named `simp.json`. Your file should like something this
```json
{
    "commands": [
        // ... you list of a commands 
        "some command",
        "another command"
    ]
}
```

> **Note:** Simp cannot execute commands that run forever e.g. nodemon, any keep-alive server you get the point

## Next Update

- 🔘 Webhook calls on successful/unsuccesful commands
- 🔘 Serve html on active server