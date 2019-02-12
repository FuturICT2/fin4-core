module Profile.Update exposing (update)

import Main.Context exposing (Context)
import Material
import Profile.Command exposing (loadProfileCmd, loadUserAssets, uploadImageCmd)
import Profile.Model exposing (Model)
import Profile.Msg exposing (Msg(..))
import Profile.Ports exposing (profileImageSelected)
import Profile.Types exposing (ViewTab(..))
import Timeline.Command
import Timeline.Model exposing (TimelineType(..))
import Timeline.Update


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    let
        userId =
            case model.data of
                Just d ->
                    d.id

                Nothing ->
                    0
    in
    case msg of
        SwitchView viewTab ->
            case model.activeViewTab == viewTab of
                True ->
                    model
                        ! [ Cmd.none ]

                False ->
                    { model | activeViewTab = viewTab }
                        ! [ case viewTab of
                                TimelineTab ->
                                    Cmd.batch
                                        [ Cmd.map TimelineMsg <| Timeline.Command.commands ctx (UserTimeline userId)
                                        ]

                                WalletTab ->
                                    Cmd.none

                                TokensTab ->
                                    loadUserAssets ctx userId 1
                          ]

        OnProfileLoadResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | data = Just data
                        , timeline = Timeline.Model.init (UserTimeline data.id)
                        , error = Nothing
                    }
                        ! [ Cmd.map TimelineMsg <| Timeline.Command.commands ctx (UserTimeline data.id) ]

                Err err ->
                    { model
                        | data = Nothing
                        , error = Just err
                    }
                        ! []

        OnUserAssetsLoadResponse resp ->
            case resp of
                Ok data ->
                    --  { model | userAssets = Just data, error = Nothing } ! []
                    let
                        loadedDataEntreis =
                            case model.userAssets of
                                Just d ->
                                    case data.page == 1 of
                                        True ->
                                            []

                                        False ->
                                            d.entries

                                Nothing ->
                                    []

                        newEntreis =
                            List.append loadedDataEntreis data.entries

                        newPage =
                            case List.length data.entries == 0 of
                                True ->
                                    data.page - 1

                                False ->
                                    data.page

                        newData =
                            { data | entries = newEntreis, page = newPage }
                    in
                    { model
                        | userAssets = Just newData
                        , isLoadingMore = False
                        , error = Nothing
                    }
                        ! []

                Err err ->
                    { model
                        | userAssets = Nothing
                        , error = Just err
                    }
                        ! []

        LoadMore ->
            let
                p =
                    case model.userAssets of
                        Just userAssets ->
                            userAssets.page

                        Nothing ->
                            1
            in
            { model | isLoadingMore = True } ! [ loadUserAssets ctx userId (p + 1) ]

        UploadProfileImage ->
            let
                base64 =
                    case model.imgInput of
                        Just d ->
                            d.contents

                        Nothing ->
                            ""
            in
            { model
                | isUploadingImage = True
                , uploadImageError = Nothing
            }
                ! [ uploadImageCmd ctx userId base64 ]

        OnUploadProfileImageResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | isUploadingImage = False
                        , uploadImageError = Nothing
                        , imgInput = Nothing
                        , data = Nothing
                    }
                        ! [ loadProfileCmd ctx userId ]

                Err err ->
                    { model
                        | isUploadingImage = False
                        , error = Just err
                    }
                        ! []

        ImageSelected ->
            model ! [ profileImageSelected "profile-img-input" ]

        ImageRead data ->
            let
                _ =
                    Debug.log "----- IMAGE READ --  ---- " msg
            in
            { model | imgInput = Just data } ! []

        TimelineMsg msg_ ->
            let
                ( childModel, cmd ) =
                    Timeline.Update.update ctx msg_ model.timeline
            in
            { model | timeline = childModel } ! [ Cmd.map TimelineMsg cmd ]

        Mdl msg_ ->
            Material.update Mdl msg_ model
