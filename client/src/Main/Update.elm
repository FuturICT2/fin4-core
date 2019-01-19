module Main.Update exposing (mountRoute, update)

import CreateToken.Model
import CreateToken.Update
import Debug
import Main.Command exposing (logoutCmd)
import Main.Context exposing (DeviceSize(..))
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        , loginPath
        , parseLocation
        , tokensPath
        )
import Material
import Navigation exposing (newUrl)
import Person.Command
import Person.Model
import Person.Update
import Portfolio.Command
import Portfolio.Model
import Portfolio.Update
import Token.Command
import Token.Model
import Token.Update
import Tokens.Command
import Tokens.Model
import Tokens.Update
import UserLogin.Update


mountRoute : Model -> ( Model, Cmd Msg )
mountRoute model =
    case model.context.route of
        TokenRoute tokenId ->
            let
                token =
                    Token.Model.init
            in
            { model | token = token } ! [ Cmd.map Token <| Token.Command.commands model.context tokenId ]

        PersonRoute personId ->
            let
                person =
                    Person.Model.init
            in
            { model | person = person } ! [ Cmd.map Person <| Person.Command.commands model.context personId ]

        PortfolioRoute ->
            model ! [ Cmd.map Portfolio <| Portfolio.Command.commands model.context ]

        TokensRoute ->
            model ! [ Cmd.map Tokens <| Tokens.Command.commands model.context ]

        _ ->
            model ! []


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    let
        _ =
            Debug.log "" msg
    in
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

        Token msg_ ->
            let
                ( childModel, cmd ) =
                    Token.Update.update model.context msg_ model.token
            in
            { model | token = childModel } ! [ Cmd.map Token cmd ]

        Person msg_ ->
            let
                ( childModel, cmd ) =
                    Person.Update.update model.context msg_ model.person
            in
            { model | person = childModel } ! [ Cmd.map Person cmd ]

        Portfolio msg_ ->
            let
                ( childModel, cmd ) =
                    Portfolio.Update.update model.context msg_ model.portfolio
            in
            { model | portfolio = childModel } ! [ Cmd.map Portfolio cmd ]

        Tokens msg_ ->
            let
                ( childModel, cmd ) =
                    Tokens.Update.update model.context msg_ model.tokens
            in
            { model | tokens = childModel } ! [ Cmd.map Tokens cmd ]

        CreateToken msg_ ->
            let
                ( childModel, cmd ) =
                    CreateToken.Update.update model.context msg_ model.createToken
            in
            { model | createToken = childModel } ! [ Cmd.map CreateToken cmd ]

        UserLoginMsg msg_ ->
            let
                ( userlogin, userloginCmd ) =
                    UserLogin.Update.update model.context msg_ model.userlogin

                context =
                    model.context

                user =
                    case userlogin.user of
                        Just _ ->
                            userlogin.user

                        _ ->
                            context.user

                cmd =
                    case userlogin.user of
                        Just _ ->
                            if userlogin.isNewUser then
                                [ newUrl "#tokens" ]

                            else
                                [ newUrl "#tokens" ]

                        _ ->
                            [ Cmd.map UserLoginMsg userloginCmd ]
            in
            { model | userlogin = userlogin, context = { context | user = user } } ! cmd

        UserLogout ->
            let
                context =
                    model.context
            in
            { model | context = { context | user = Nothing } } ! [ logoutCmd context model ]

        OnLogoutResponse resp ->
            case resp of
                Ok _ ->
                    let
                        context =
                            model.context

                        cmd =
                            case context.user of
                                Just _ ->
                                    []

                                _ ->
                                    [ newUrl loginPath ]
                    in
                    { model | context = { context | user = Nothing } } ! cmd

                Err _ ->
                    model ! []

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
