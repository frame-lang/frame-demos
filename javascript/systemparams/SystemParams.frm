#SystemParams $[stateMsg:string] >[enterMsg:string]

    -machine-

    $Begin [stateMsg:string]
        |>|[enterMsg:string]
            print(stateMsg + " " + enterMsg) ^

    -actions-

    print[msg:string] {`
        console.log(msg)
    `}

##