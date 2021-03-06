// @flow
import {OS_ELECTRON} from './platform.shared'
import path from 'path'
import getenv from 'getenv'

const OS = OS_ELECTRON
const isMobile = false

const runMode = getenv('KEYBASE_RUN_MODE', 'prod')

if (__DEV__) {
  console.log(`Run mode: ${runMode}`)
}

const envedPathLinux = {
  staging: 'keybase.staging',
  devel: 'keybase.devel',
  prod: 'keybase',
}

const envedPathOSX = {
  staging: 'KeybaseStaging',
  devel: 'KeybaseDevel',
  prod: 'Keybase',
}

const envedPathWin32 = {
  staging: 'KeybaseStaging',
  devel: 'KeybaseDevel',
  prod: 'Keybase',
}

function buildWin32SocketRoot () {
  let appdata = getenv('APPDATA', '')
  // Remove leading drive letter e.g. C:
  if (/^[a-zA-Z]:/.test(appdata)) {
    appdata = appdata.slice(2)
  }
  // Handle runModes, prod has no extension.
  let extension = ''
  if (runMode !== 'prod') {
    extension = runMode.charAt(0).toUpperCase() + runMode.substr(1)
  }
  let path = `\\\\.\\pipe\\kbservice${appdata}\\Keybase${extension}`
  return path
}

function findSocketRoot () {
  const paths = {
    'darwin': `${getenv('HOME', '')}/Library/Caches/${envedPathOSX[runMode]}/`,
    'linux': runMode === 'prod' ? `${getenv('XDG_RUNTIME_DIR', '')}/keybase/` : `${getenv('XDG_RUNTIME_DIR', '')}/keybase.${runMode}/`,
    'win32': buildWin32SocketRoot(),
  }

  return paths[process.platform]
}

function findDataRoot () {
  const linuxDefaultRoot = `${getenv('HOME', '')}/.local/share`
  const paths = {
    'darwin': `${getenv('HOME', '')}/Library/Application Support/${envedPathOSX[runMode]}/`,
    'linux': `${getenv('XDG_DATA_HOME', linuxDefaultRoot)}/${envedPathLinux[runMode]}/`,
    'win32': `${getenv('APPDATA', '')}\\Keybase\\`,
  }

  return paths[process.platform]
}

function logFileName () {
  const paths = {
    'darwin': `${getenv('HOME', '')}/Library/Logs/${envedPathOSX[runMode]}.app.log`,
    'linux': null, // linux is null because we can redirect stdout
    'win32': `${getenv('APPDATA', '')}\\${envedPathWin32[runMode]}\\keybase.app.log`,
  }

  return paths[process.platform]
}

const socketRoot = findSocketRoot()
const socketName = 'keybased.sock'
const socketPath = path.join(socketRoot, socketName)
const dataRoot = findDataRoot()
const splashRoot = process.platform === 'darwin' ? socketRoot : dataRoot

export {
  OS,
  dataRoot,
  isMobile,
  logFileName,
  runMode,
  socketName,
  socketPath,
  socketRoot,
  splashRoot,
}
