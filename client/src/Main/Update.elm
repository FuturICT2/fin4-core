module Main.Update exposing (mountRoute, update)

import CreateToken.Model
import CreateToken.Update
import Debug
import Main.Context exposing (DeviceSize(..))
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        , parseLocation
        )
import Material
import Navigation exposing (newUrl)
import Portfolio.Command
import Portfolio.Model
import Portfolio.Update
import Tokens.Command
import Tokens.Model
import Tokens.Update


mountRoute : Model -> ( Model, Cmd Msg )
mountRoute model =
    case model.context.route of
        TokensRoute ->
            model ! [ Cmd.map Tokens <| Tokens.Command.commands model.context ]

        PortfolioRoute ->
            model ! [ Cmd.map Portfolio <| Portfolio.Command.commands model.context ]

        _ ->
            model ! []


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        OnRouteChange location ->
            let
                context =
                    model.context
            in
            mountRoute
                { model
                    | context =
                        { context
                            | route = parseLocation location
                        }
                    , showMobileNav = False
                    , createToken = CreateToken.Model.init
                }

        OnCheckSessionResponse resp ->
            let
                context =
                    model.context
            in
            case resp of
                Ok user ->
                    { model | context = { context | user = Just user, sessionDidLoad = True } } ! []

                Err _ ->
                    { model | context = { context | sessionDidLoad = True } } ! []

        Homepage msg_ ->
            model ! []

        Tokens msg_ ->
            let
                ( childModel, cmd ) =
                    Tokens.Update.update model.context msg_ model.tokens
            in
            { model | tokens = childModel } ! [ Cmd.map Tokens cmd ]

        Portfolio msg_ ->
            let
                ( childModel, cmd ) =
                    Portfolio.Update.update model.context msg_ model.portfolio
            in
            { model | portfolio = childModel } ! [ Cmd.map Portfolio cmd ]

        CreateToken msg_ ->
            let
                ( childModel, cmd ) =
                    CreateToken.Update.update model.context msg_ model.createToken
            in
            { model | createToken = childModel } ! [ Cmd.map CreateToken cmd ]

        OnWindowResize size ->
            let
                device =
                    if size.width <= 950 then
                        Mobile

                    else
                        Desktop

                window =
                    { width = size.width
                    , height = size.height
                    , device = device
                    , isMobile = device == Mobile
                    }

                context =
                    model.context
            in
            { model | context = { context | window = window } } ! []

        ToggleMobileNav ->
            { model | showMobileNav = not model.showMobileNav } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
