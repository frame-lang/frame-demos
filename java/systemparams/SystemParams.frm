#SystemParams $[stateMsg:String] >[enterMsg:String]

    -machine-

    $Begin [stateMsg:String]
        |>|[enterMsg:String]
            print(stateMsg + " " + enterMsg) ^

    -actions-

    print[msg:String] {`
        console.log(msg)
    `}

##