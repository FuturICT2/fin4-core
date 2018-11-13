'use strict'

import './res/style.scss'
require('./index.html');
const Elm = require('./Main.elm')
const apiBase = '/wapi'
const Identicon = require('identicon.js')

const app = Elm.Main.embed(document.getElementById('app'), {
  apiBase: apiBase
});
