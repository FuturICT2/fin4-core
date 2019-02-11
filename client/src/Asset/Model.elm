module Asset.Model exposing (Asset, Block, Image, Miner, Model, assetDecoder, assetListDecoder, init)

import Http
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Material
import Timeline.Model exposing (TimelineType(..))


type alias Block =
    { id : Int
    , userId : Int
    , userName : String
    , text : String
    , status : Int
    , ytVideoId : String
    , createdAt : String
    , createdAtHuman : String
    , images : List String
    }


type alias Miner =
    { id : Int
    , userName : String
    , mined : String
    , miningPercentage : String
    }


type alias Asset =
    { id : Int
    , name : String
    , symbol : String
    , creatorId : Int
    , creatorName : String
    , description : String
    , totalSupply : Int
    , minersCount : Int
    , favoritesCount : Int
    , didUserLike : Bool
    , miners : List Miner
    }


type alias Image =
    { contents : String
    , filename : String
    }


type alias Model =
    { mdl : Material.Model
    , data : Maybe Asset
    , tab : Int
    , imgs : List Image
    , blockText : String
    , isSubmittingBlock : Bool
    , submitBlockError : Maybe Http.Error
    , error : Maybe String
    , timeline : Timeline.Model.Model
    }


init : Model
init =
    { mdl = Material.model
    , data = Nothing
    , tab = 0
    , imgs = []
    , blockText = ""
    , isSubmittingBlock = False
    , submitBlockError = Nothing
    , error = Nothing
    , timeline = Timeline.Model.init (AssetTimeline 0)
    }


assetDecoder : JD.Decoder Asset
assetDecoder =
    JP.decode Asset
        |> JP.required "ID" JD.int
        |> JP.required "Name" JD.string
        |> JP.required "Symbol" JD.string
        |> JP.required "CreatorID" JD.int
        |> JP.required "CreatorName" JD.string
        |> JP.required "Description" JD.string
        |> JP.required "Supply" JD.int
        |> JP.required "MinersCounter" JD.int
        |> JP.required "FavoritesCounter" JD.int
        |> JP.required "DidUserLike" JD.bool
        |> JP.optional "Miners" (JD.list minerDecoder) []


assetListDecoder : JD.Decoder (List Asset)
assetListDecoder =
    JD.list assetDecoder


blockDecoder : JD.Decoder Block
blockDecoder =
    JP.decode Block
        |> JP.required "ID" JD.int
        |> JP.required "UserID" JD.int
        |> JP.required "UserName" JD.string
        |> JP.required "Text" JD.string
        |> JP.required "Status" JD.int
        |> JP.required "YtVideoID" JD.string
        |> JP.required "CreatedAt" JD.string
        |> JP.required "CreatedAtHuman" JD.string
        |> JP.optional "Images" (JD.list JD.string) []


minerDecoder : JD.Decoder Miner
minerDecoder =
    JP.decode Miner
        |> JP.required "ID" JD.int
        |> JP.required "UserName" JD.string
        |> JP.required "Mined" JD.string
        |> JP.required "MiningPercentage" JD.string
