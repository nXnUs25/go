# Questions 3
---
## Questions

1. Why is software versioning important?
2. What does a semantic version look like and when are its individual components incremented?
3. Which component of a package's semantic version would you increment in the following cases?
    - A new API is introduced.
    - An existing API is modified and a new, required parameter is added to it.
    - A fix for a security bug is committed.

4. Name some alternative versioning schemes that we could use besides semantic versioning.
5. What are the pros and cons of vendoring?
6. Name some of the differences between the dep tool and Go modules.”

> Excerpt From
> Hands-On Software Engineering with Golang
> Achilleas Anagnostopoulos
> This material may be protected by copyright.

---

## Answers 3

### 3.1
Nowadays, as more and more software engineers ascribe to the release fast mantra, having a sane versioning system in place makes it possible to do the following:

- Validate that a particular piece of software can be used as a safe drop-in replacement for an older piece of software that we are using as part of our production systems. This is especially important from a security standpoint, as all software, unless formally verified, may contain potential security bugs that can be discovered at any point in time—even weeks or years after it has been deployed to production. It is therefore of paramount importance for us to be able to mitigate such issues by upgrading to a newer release as soon as bug fixes become available.
- Pin down each dependency of our applications to a particular package version. This is a key prerequisite for setting up CI pipelines to implement the concept of repeatable builds. Having access to repeatable builds makes it possible to recompile, at any time, an exact copy of the software that a customer runs in production and use it as a reference when investigating bug reports.

---
### 3.2
**Semantic versioning** is a widely popular system for describing software versions in a way that makes it quite straightforward for the intended software users to figure out which versions are safe to upgrade to and which versions contain breaking API changes and therefore require development effort and time when upgrading.

Semantic versions are formatted as follows:
**`MAJOR.MINOR.PATCH`**

Depending on the use case, additional suffixes may be optionally appended to indicate a prerelease (for example, alpha, beta, or RC (or release candidate)) or to convey other build-related information (for example, the Git SHA for the branch that is used to build a release or perhaps a timestamp for when the build artifacts were generated).

**Semantic versioning** makes it easy for package authors to let users of the package know what type of changes each release contains. 
For example, 
- the **_`PATCH`_** field is incremented whenever a *`backward-compatible`* bug fix is applied to the code. 
- the **_`MINOR`_** field is incremented when *`new functionality gets added to a package`*, but, most importantly, only when this new functionality is added in a manner that ensures that the **`new version remains backward compatible`** with older package versions. 
- the **_`MAJOR`_** component of the version string will need to be incremented. as packages evolve over time, it is inevitable that at some point some breaking changes will have to be introduced. For instance, an existing function signature may need to be changed to support additional use cases.

---
### 3.3
    
- MAJOR - 0.1.0 for _"A new API is introduced."_
- MINOR - 1.1.0 for _"An existing API is modified and a new, required parameter is added to it."_
- PATCH - 1.1.1 for _"A fix for a security bug is committed."_

> _**Note****_ had to correct above since i had different answer than was in book .. before i had:
> - MAJOR - 1.0.0 for _"A new API is introduced."_
> - MINOR - 1.1.0 for _"An existing API is modified and a new, required parameter is added to it."_
>   since i ws thinking that we introducing a new API

---
### 3.4
- Many projects use a date-based versioning scheme called **Calendar Versioning** (aka *CalVer*) as example Ubuntu
  When using dates in versioning, for instance, file names, it is common to use the **ISO 8601** scheme `YYYY-MM-DD`, as this is easily string-sorted in increasing or decreasing order. The hyphens are sometimes omitted.
- Python **PEP 440** – Version Identification and Dependency Specification, outlining their own flexible scheme, that defines an epoch segment, a release segment, pre-release and post-release segments and a development release segment.
- **`BLAG Linux and GNU`** features very large version numbers: _major_ releases have numbers such as _50000 and 60000_, while _minor_ releases increase the number by 1 (e.g. 50001, 50002). Alpha and beta releases are given decimal version numbers slightly less than the major release number, such as 19999.00071 for alpha 1 of version 20000, and 29999.50000 for beta 2 of version 30000. Starting at 9001 in 2003, the most recent version as of 2011 is 140000.
- **Urbit** uses Kelvin versioning (named after the absolute Kelvin temperature scale): software versions start at a high number and count down to version 0, at which point the software is considered finished and no further modifications are made.

---
### 3.5
- Benefits of vendoring dependencies:

  * First and foremost, the key promise of vendoring is nothing other than the capability to run reproducible builds. Many customers, especially larger corporations, tend to stick to stable or LTS releases for the software they deploy and forego upgrading their systems unless it's absolutely necessary. Being able to check out the exact software version that a customer uses and generate a bit-for-bit identical binary for use in a test environment is an invaluable tool for any field engineer attempting to diagnose and reproduce bugs that the customers are facing.

  * Another benefit of vendoring is that it serves as a safety net in case an upstream dependency suddenly disappears from the place where it is hosted (for example, a GitHub or GitLab repository), thereby breaking builds for software that depends on it. If you are thinking that this is a highly unlikely scenario, let me take you back to 2016 and share an interesting engineering horror story from the world of Node.js!

- Cons of vendoring dependencies:
  - You may have heard of the now-infamous left-pad package. In case you haven't, it is just a single-function package that, as you can probably figure out by its name, provides a function to pad a string up to a specific length with a specific character. Nothing really scary so far... except that this small package was a direct dependency of over 500 packages, which in turn were transient dependencies of several other packages, and so on. Everything went fine until one day, the left-pad package maintainer received a cease-and-desist letter for one of his other packages and decided, as a form of protest, to take down all of his packages, including left-pad.

  - Now, picture the chaos that ensued as peoples CI builds started breaking one after the other. But engineering teams that had been judiciously vendoring their dependencies were not affected by this issue at all.

---
### 3.6
Some differences between the dep tool and Go modules are as follows: 
   - Go modules fully integrate with the various commands, such as `go get`, `go build`, and `go test`. 
   - While the dep tool selects the **highest** common version for a package, 
   - Go modules select the **minimum** viable version. Go modules support multi-versioned dependencies. 
   - Go modules do away with the vendor folder that's used by the dep tool.

---

> Excerpt From
> Hands-On Software Engineering with Golang
> Achilleas Anagnostopoulos
> This material may be protected by copyright.

---

# Answers to Chapter 3 ... by author of the book 

### Chapter 3 

#### 1. 
The purpose of software versioning is twofold. First, it allows software engineers to validate whether an external dependency can be safely upgraded without the risk of introducing issues to production systems. Secondly, being able to explicitly reference required software dependencies via their versions is a prerequisite for implementing the concept of repeatable builds. 
   
#### 2. 
A semantic version is a string that satisfies the following format: 
   `MAJOR.MINOR.PATCH`: 
      - The major component is incremented when a breaking change is introduced to the software 
      - The minor component is incremented when new functionality is introduced to the software in a backward-compatible way 
      - The patch version is incremented when a backward-compatible fix is applied to the code 

#### 3. 
We would need to: 
   - In the first case, we would increment the minor version as the new API does not break backward compatibility. 
   - In the second case, we would increment the major version as the new required parameter breaks compatibility with older versions of the API. 
   - Finally, in the third case, we would increment the patch version. 

#### 4. 
One approach would be to tag each build with a unique, monotonically increasing number. Alternatively, we could annotate build artifacts with a timestamp that indicates when they were created. 

#### 5. 
The pros of vendoring are as follows: 
   - The capability to run reproducible builds for current or older versions of a piece of software 
   - Being able to access the required dependencies locally, even if they disappear from the place where they were originally hosted

The cons of vendoring are as follows: 
   - Engineers should monitor the change logs for their dependencies and manually upgrade them when security fixes become available. If the authors of the vendored dependencies do not follow semantic versioning for their packages, upgrading a dependency can introduce breaking changes that must be addressed before it's able to compile our code. 

#### 6. 
Some differences between the dep tool and Go modules are as follows: 
   - Go modules fully integrate with the various commands, such as `go get`, `go build`, and `go test`. 
   - While the dep tool selects the **highest** common version for a package, 
   - Go modules select the **minimum** viable version. Go modules support multi-versioned dependencies. 
   - Go modules do away with the vendor folder that's used by the dep tool.
