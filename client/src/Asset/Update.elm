module Asset.Update exposing (update)

import Asset.Command exposing (loadAssetCmd, submitBlockCmd, verifyBlockCmd)
import Asset.Model exposing (Model)
import Asset.Msg exposing (Msg(..))
import Asset.Ports exposing (assetBlockImageSelected)
import Dict exposing (Dict)
import Gallery
import Main.Context exposing (Context)
import Material
import Timeline.Command
import Timeline.Model exposing (TimelineType(..))
import Timeline.Update


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    let
        getAssetId m =
            case model.data of
                Just v ->
                    v.id

                Nothing ->
                    0
    in
    case msg of
        OnAssetLoadResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | data = Just data
                        , timeline = Timeline.Model.init (AssetTimeline data.id)
                    }
                        ! [ Cmd.map TimelineMsg <| Timeline.Command.commands ctx (AssetTimeline data.id) ]

                Err _ ->
                    { model
                        | data = Nothing
                        , error = Just "Error loading token. Please try again."
                    }
                        ! []

        SelectTab tab ->
            { model | tab = tab } ! []

        SetBlockText value ->
            { model | blockText = value } ! []

        SubmitBlock ->
            let
                assetId =
                    getAssetId model
            in
            { model
                | isSubmittingBlock = True
                , submitBlockError = Nothing
            }
                ! [ submitBlockCmd ctx assetId model.blockText model.imgs ]

        SubmitBlockResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | blockText = ""
                        , isSubmittingBlock = False
                        , submitBlockError = Nothing
                        , imgs = []
                    }
                        ! [ loadAssetCmd ctx (getAssetId model) ]

                Err error ->
                    { model
                        | isSubmittingBlock = False
                        , submitBlockError = Just error
                    }
                        ! []

        VerifyBlock isRejected blockId ->
            model ! [ verifyBlockCmd ctx isRejected blockId ]

        VerifyBlockResponse resp ->
            case resp of
                Ok data ->
                    model ! [ loadAssetCmd ctx (getAssetId model) ]

                Err error ->
                    model ! []

        ImageSelected ->
            model ! [ assetBlockImageSelected <| "block-img" ]

        ImageRead data ->
            { model | imgs = List.append model.imgs [ data ] } ! []

        DeleteImage idx ->
            let
                removeFromList i xs =
                    List.take i xs ++ List.drop (i + 1) xs
            in
            { model | imgs = removeFromList idx model.imgs } ! []

        TimelineMsg msg_ ->
            let
                ( childModel, cmd ) =
                    Timeline.Update.update ctx msg_ model.timeline
            in
            { model | timeline = childModel } ! [ Cmd.map TimelineMsg cmd ]

        Mdl msg_ ->
            Material.update Mdl msg_ model
