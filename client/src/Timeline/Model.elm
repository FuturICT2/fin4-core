module Timeline.Model exposing
    ( Model
    , TimelineEntries
    , TimelineEntry
    , TimelineType(..)
    , init
    , timelineEntriesDecoder
    )

import Dict exposing (Dict)
import Gallery
import Http
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Material


type TimelineType
    = UserTimeline Int
    | HomepageTimeline
    | AssetTimeline Int


type alias TimelineEntry =
    { blockId : Int
    , userId : Int
    , userName : String
    , userProfileImageURL : String
    , assetId : Int
    , assetName : String
    , assetSymbol : String
    , oracleId : Int
    , oracleName : String
    , text : String
    , status : Int
    , ytVideoId : String
    , favoritesCount : Int
    , didUserLike : Bool
    , createdAt : String
    , createdAtHuman : String
    , images : List String
    }


type alias TimelineEntries =
    { hasMore : Bool
    , page : Int
    , entries : List TimelineEntry
    }


type alias Model =
    { mdl : Material.Model
    , timelineType : TimelineType
    , data : Maybe TimelineEntries
    , gallery : Dict Int Gallery.State
    , isLoadingMore : Bool
    , loadMoreError : Maybe Http.Error
    , error : Maybe Http.Error
    }


init : TimelineType -> Model
init timelineType =
    { mdl = Material.model
    , timelineType = timelineType
    , data = Nothing
    , gallery = Dict.empty
    , isLoadingMore = False
    , loadMoreError = Nothing
    , error = Nothing
    }


timelineEntriesDecoder : JD.Decoder TimelineEntries
timelineEntriesDecoder =
    JP.decode TimelineEntries
        |> JP.required "HasMore" JD.bool
        |> JP.required "Page" JD.int
        |> JP.required "Entries" (JD.list timelineEntryDecoder)


timelineEntryDecoder : JD.Decoder TimelineEntry
timelineEntryDecoder =
    JP.decode TimelineEntry
        |> JP.required "BlockID" JD.int
        |> JP.required "UserID" JD.int
        |> JP.required "UserName" JD.string
        |> JP.required "UserProfileImageURL" JD.string
        |> JP.required "AssetID" JD.int
        |> JP.required "AssetName" JD.string
        |> JP.required "AssetSymbol" JD.string
        |> JP.required "OracleID" JD.int
        |> JP.required "OracleName" JD.string
        |> JP.required "Text" JD.string
        |> JP.required "Status" JD.int
        |> JP.required "YtVideoID" JD.string
        |> JP.required "FavoritesCount" JD.int
        |> JP.required "DidUserLike" JD.bool
        |> JP.required "CreatedAt" JD.string
        |> JP.required "CreatedAtHuman" JD.string
        |> JP.optional "Images" (JD.list JD.string) []
