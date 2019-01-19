module Model.Persons exposing (MinedToken, Person, personDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP


type alias MinedToken =
    { balance : String
    , mined : String
    , tokenId : Int
    , tokenName : String
    , tokenSymbol : String
    }


type alias Person =
    { id : Int
    , name : String
    , email : String
    , minedTokens : List MinedToken
    }


personDecoder : JD.Decoder Person
personDecoder =
    JP.decode Person
        |> JP.required "ID" JD.int
        |> JP.required "Name" JD.string
        |> JP.required "Email" JD.string
        |> JP.required "MinedTokens" (JD.list tokenDecoder)


tokenDecoder : JD.Decoder MinedToken
tokenDecoder =
    JP.decode MinedToken
        |> JP.required "Balance" JD.string
        |> JP.required "Mined" JD.string
        |> JP.required "TokenID" JD.int
        |> JP.required "TokenName" JD.string
        |> JP.required "TokenSymbol" JD.string
