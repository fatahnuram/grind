# grind

Helps you track your grind! :muscle:

List out your grind daily in the local environment via CLI. Data storage relies on local files (json and csv files), thus doesn't need any internet connection (can be run on an env without any internet connectivity).

## Installation

Currently we support installation from source. Other type of installation might be supported in the future.

### From source

1. Make sure you have Golang in your local workspace (1.21+), or [download it here](https://go.dev/doc/install).

2. Make sure to update your `$PATH` to include your `$GOBIN`.

3. Run `go install`.

4. Let's go `grind`-ing! :muscle:

## Introduction

We all have at least one activity (or more), which may become a routine already. Some of them might be important for our personal life goal and growth (which we might not aware). Maybe it's not about personal goal, maybe it's about business goal, or any other goals you want to pursue. But on the other side, we have a very limited time. That's why we need to focus! :dart:

`grind` helps you to track down your daily, weekly, or even monthly grind activities. It's pretty similar with to-do list apps out there, but the difference is that `grind` is mainly focusing for your recurring activities (or routine, you say).

## How to use

You can follow these simplified guide:

1. Download and install `grind` to your local workspace.

2. Prepare your grind activities. You can check out the `template.csv` file as a starting point (by renaming it to `grind.csv` to make it work).

3. Run the `grind` at least once daily. Try to complete all the listed items there! :100:

## Template file

A `template.csv` file can be found in this project as a template file and may be used as an example to build upon. Copy the template and rename it to `grind.csv` so that the app will use it as the main data source.

Supported fields (columns) are: `name`, `frequency`, `day`, `date`, `month`, `function`. First row will always be treated as header, means that the first activity/item is always in the second row.

Along the journey, the app will produce and operates on a file called `progress.json` to track down the progress.

**NOTE**: removing `grind.csv` means you wipe out your only data source, while removing `progress.json` means you clear out your whole progress (restarting from zero).

### name

`name` is the name of the activity.

### frequency

`frequency` is the frequency of the activity.

Possible values are: `daily`, `weekly`, `monthly`, `annually`, or `custom`.

Each acitivities will need certain columns to be non-empty, otherwise it may not function properly. At the very least, the column `name` and `frequency` are always required.

Required fields for each values (while others *will be ignored*):

- `daily` doesn't require any other fields
- `weekly` requires `day`
- `monthly` requires `date`
- `annually` requires both `date` and `month`
- `custom` requires `function`

### day

`day` is to specify what day the activity starts to be available. Only valid when `frequency` is `weekly`.

Possible values are: `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.

Number representation of `sun` is `1`, while `sat` is `7`.

**NOTE**: For `day` values, it is required to use one of the possible values above **and not using** the number representation. These number representations won't be used in the `day` field, but instead it will be used in the `function` field below.

### date

`date` is to specify what date the activity starts to be available. Only valid when `frequency` is either `monthly` or `annually`.

Possible values are: `1-31` with some special exceptions (e.g. `feb` will only have `1-28` or `1-29`, while `apr` will only have `1-30`).

### month

`month` is to specify what month the activity starts to be available. Only valid when the `frequency` is `annually` (combined with the `date` field).

Possible values are: `jan`, `feb`, `mar`, `apr`, `may`, `jun`, `jul`, `aug`, `sep`, `oct`, `nov`, `dec`.

Number representation of `jan` is `1`, while `dec` is `12`.

**NOTE**: For `month` values, it is required to use one of the possible values above **and not using** the number representation. These number representations won't be used in the `month` field, but instead it will be used in the `function` field below.

### function

`function` is to specify when the activity starts to be available when the `frequency` is `custom`.

The app will use this custom function to define if an activity is valid (i.e. available) or not. A certain activity will be valid/available when the function evaluates to `true`.

You can use `day`, `date`, `month`, and `year` as built-in variables to define the activity's validity function.

Some examples:

- `date%3==0` means if the date is divisible by 3 then the activity is available, i.e. `3`, `6`, etc
- `day==3 || day==5` means if the day's number representation is 3 or 5 then the activity is available, i.e. `tue` or `thu`
- `year%5==0 && month==2` means if the year is divisible by 5 and the month's number representation is 2 then the activity is available, i.e. `2020 feb`, `2025 feb`, etc

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE)
