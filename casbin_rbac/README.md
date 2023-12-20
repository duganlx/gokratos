# Casbin RBAC

## Entity Type

user: Authorized Entity

userGroup: User Group

srcGroup: Resources Group

src: Resources


## Entity Association Relationship

`user -> userGroup -> srcGroup -> src`

All relationships will ultimately be transformed into relationships between 'user' and 'src'. 'userGroup' and 'srcGroup' are just for the convenience of implementing the relationship between 'user' and 'src'. So, 'userGroup' and 'srcGroup' are not necessarily required.

Therefore, I recommend considering 'userGroup' and 'srcGroup' as different dimensions of resource management. There are three levels in total: the lowest level is 'src,' representing a single resource; the next level is 'srcGroup,' representing a group of resources; the highest level is 'userGroup,' which can include both 'srcGroup' and 'src.' It is important to note that at any level, the ultimate reference is to specific resources, namely 'src'. All scenarios are as follows.

```text
user -> src
user -> userGroup -> src
user -> userGroup -> srcGroup -> src
userGroup -> src
userGroup -> srcGroup -> src
srcGroup -> src 
```

## Scenario Setup

### Entity Object Relationship

user: lvx, xjw, ww

userGroup: admin, quant, it

srcGroup: eam, drw, test, prod, v1, v2

src: The naming convention is `<eam/drw>0<1-5>:<v1/v2>:<prod/test>`, where `<>` indicates making a choice within.


### Entities Resource Relationship

`USER:lvx -> [USERGROUP:it, SRCGROUP:test, drw05:v1:test]`: The user 'lvx' has access to 'userGroup:it', 'srcGroup:test' and 'drw05:v1:test'.

`USER:xjw -> [USERGROUP:admin, USERGROUP:it]`

`USER:ww -> [USERGROUP:quant]`

`USERGROUP:it -> [eam0<1-3>:v1:prod, <*>0<*>:v1:test]`

`USERGROUP:quant -> <eam/drw>0<*>:v2:prod`

`USERGROUP:admin -> <*>0<*>:<*>:prod`

`SRCGROUP:eam ->`

`SRCGROUP:drw ->`

`SRCGROUP:test ->`

`SRCGROUP:prod ->`

`SRCGROUP:v1 ->`

`SRCGROUP:v2 ->`

Note:
1. To facilitate differentiation between different entities, a prefix will be added to objects indicating their associated entity.
2. `*` represents a wildcard.