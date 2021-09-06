# Release Guide

This guide is aimed at Release Managers tasked with cutting a new release of the Karavel Container Platform.
It should be treated as a reference guide and as a checklist to ensure that the release process goes smoothly.

## Versioning

The Karavel Container Platform is versioned using the [CalVer](https://calver.org) format due to its fixed
release schedule and rolling versioning nature. The version number is composed of a Major component (in Full Year format)
indicating the year of release, a Minor component that is incremented in each release of the same year, and a Patch component
that is incremented on every fix applied to a specific Minor release. The full format is `YYYY.X.Z`.
For example, the `2021.2.1` is the first patch of the second release of the year 2021.

## Pre-work

When the moment comes to prepare a new release, [a tracking issue should be created](https://github.com/karavel-io/platform/issues/new?labels=release).
The issue will track the release process and will host discussions and comments regarding the release.

A member of the Release Team (@projectkaravel/platform-release) will be assigned to the issue. They will be called the
"assigned Release Manager" from now one.

The issue will also be used to collect the version for each component included in the release. Component owners
will submit the proposed version by commenting on the issue. The assigned Release Manager will collect those versions and
use them to compile the release index.

## Release Candidate

Before publishing the official stable release, one or more Release Candidates (RC) will be published to identify and fix
last minute bugs in the selected components and verify that everything is ready. The RC will be tagged with a release number
formatted as `YYYY.X-RC.Z`, where `Z` is an incremental number identifying the specific Release Candidate for that version.
Notice that the Patch component is missing in the main version number. The actual release process is the same for Release Candidates
and final releases. If a problem is found and fixed in a Release Candidate, a new RC should be released with the `Z` component
incremented by 1. E.g. `YYYY.X-RC.2` updates `YYYY.X-RC.1`.

When a Release Candidate is deemed to be ready and stable, the final release will be prepared by dropping the `-RC.Z` suffix. 

## Preparing a release

When the full list of component versions has been compiled, the assigned Release Manager will create a branch from the
`main` branch called `release/yyyy.x` (or `release/yyyy.x-rc.z` for RCs) and open a Pull Request with the following changes:

- create a new folder inside [releases](releases) called `yyyy.x`, copying the content of the [template](releases/template)
folder as a starting point.
- compile a list of components to include in the release. This list is in the format `component-slug: component-version`.
- compile the CHANGELOG file with the list of components and versions in the `Added` section, linking to each component's
`CHANGELOG.md` entry for that specific version.

The Pull Request will be reviewed and approved by one or more other members of the Release Management team.

## Publishing a release

Once it is done, [a new GitHub release](https://github.com/karavel-io/platform/releases/new) should be created with
the corresponding tag called `yyyy.x.0` (notice the Patch component is now present and set a 0).
The description should be filled by copying the corresponding CHANGELOG.md section.

An automation will be triggered by the newly created release that will assemble and publish the KCP release on the `gh-pages`
branch, making it available for public download.

## Patching a release

When a patch should be issued to an existing release, a Pull Request should be opened with the appropriate changes to the 
corresponding [releases](releases) folder:
- updated component list
- new CHANGELOG.md section with the `Z` version component incremented by 1, and appropriate entries describing what has been
changed and/or fixed.
  
The corresponding release will then be created on [GitHub](https://github.com/karavel-io/platform/releases/new) alongside
the corresponding `YYYY.X.Z` tag.
The description should be filled by copying the corresponding CHANGELOG.md section.

An automation will be triggered by the newly created release that will assemble and publish the KCP release on the `gh-pages`
branch, making it available for public download.
