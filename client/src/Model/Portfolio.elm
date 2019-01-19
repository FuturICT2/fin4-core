module Model.Portfolio exposing (Portfolio, Position, init, portfolioDecoder, positionDecoder)

import Json.Decode as JD
import Json.Decode.Pipeline as JP


type alias Position =
    { userId : Int
    , tokenId : Int
    , balance : String
    , reserved : String
    , tokenName : String
    , tokenSymbol : String
    , logoFile : String
    , blockchainAddress : String
    }


type alias Portfolio =
    { positions : List Position
    }


init : Portfolio
init =
    { positions = []
    }


portfolioDecoder : JD.Decoder Portfolio
portfolioDecoder =
    JP.decode Portfolio
        |> JP.required "Entries" (JD.list positionDecoder)


positionDecoder : JD.Decoder Position
positionDecoder =
    JP.decode Position
        |> JP.required "UserID" JD.int
        |> JP.required "TokenID" JD.int
        |> JP.required "Balance" JD.string
        |> JP.required "Reserved" JD.string
        |> JP.required "TokenName" JD.string
        |> JP.required "TokenSymbol" JD.string
        |> JP.required "LogoFile" JD.string
        |> JP.required "BlockchainAddress" JD.string
