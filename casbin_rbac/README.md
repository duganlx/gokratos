# Casbin RBAC

## Entity Type

user: Authorized Entity

userGroup: User Group

srcGroup: Resources Group

src: Resources

## Entity Association Relationship

`user -> userGroup -> srcGroup -> src`

All relationships will ultimately be transformed into relationships between 'user' and 'src'. 'userGroup' and 'srcGroup' are just for the convenience of implementing the relationship between 'user' and 'src'. So, 'userGroup' and 'srcGroup' are not necessarily required.

Therefore, I recommend considering 'userGroup' and 'srcGroup' as different dimensions of resource management. There are three levels in total: the lowest level is 'src', representing a single resource; the next level is 'srcGroup', representing a group of resources; the highest level is 'userGroup', which can include both 'srcGroup' and 'src'. It is important to note that at any level, the ultimate reference is to specific resources, namely 'src'. All scenarios are as follows.

```text
user -> src
user -> userGroup -> src
user -> userGroup -> srcGroup -> src
userGroup -> src
userGroup -> srcGroup -> src
srcGroup -> src 
```

## Scenario Setup

为了能够尽可能的考虑到所有的权限管理情况，我们将模拟一个虚拟的基金公司，并对它进行权限管理。假设该公司有三个部门，分别是 信息技术部(IT)、量化策略部(quant)、市场运营部(operation)。该公司会发行两个系列的产品，分别叫做 EAM 和 DRW，同时，每个系列都会有两个产品，代号为01 和 02。这些产品也会进行版本迭代，目前存在两个版本 v1 和 v2，v2 是当前最新版本。这些产品实际上是一个服务，对此有两种操作权限，分别是 读r 和 写w。并且为了能够测试，所有的服务都会运行在两个环境中，分别是测试环境 test 和生产环境 prod。

### Entity Object Relationship

user: lvx, xjw, ww, wyp, dxs

userGroup: admin(xjw), quant(ww), it(lvx), operation(dxs), mgr(wyp)

srcGroup: eam, drw, test, prod, v1, v2, ...(And the combinations between them.)

src: The naming convention is `<eam/drw><01/02>:<v1/v2>:<prod/test>`, where `<>` indicates making a choice within.

privilege: ro, wo, rw

### Entities Resource Relationship

资源组和资源之间的关系如下所示。需要注意的是，资源组不支持嵌套，即没有“子资源组”的概念，所以资源组 eam 和 eamv1test 虽然看上去存在包含关系，但实际它们属于同一级别的。后续需要做成可支持子资源组。

```
SRCGROUP:eam -> [SRCGROUP:eamv1test, SRCGROUP:eamv1prod, SRCGROUP:eamv2test, SRCGROUP:eamv2prod] -> [eam<01,02>:<v1,v2>:<prod,test>]
SRCGROUP:drw -> [SRCGROUP:drwv1test, SRCGROUP:drwv1prod, SRCGROUP:drwv2test, SRCGROUP:drwv2prod] -> [drw<01,02>:<v1,v2>:<prod,test>]
SRCGROUP:v1 -> [SRCGROUP:eamv1test, SRCGROUP:eamv1prod, SRCGROUP:drwv1test, SRCGROUP:drwv1prod] -> [<eam,drw><01,02>:v1:<prod,test>]
SRCGROUP:v2 -> [SRCGROUP:eamv2test, SRCGROUP:eamv2prod, SRCGROUP:drwv2test, SRCGROUP:drwv2prod] -> [<eam,drw><01,02>:v2:<prod,test>]
SRCGROUP:test -> [SRCGROUP:eamv1test, SRCGROUP:eamv2test, SRCGROUP:drwv1test, SRCGROUP:drwv2test] -> [<eam,drw><01,02>:<v1,v2>:test]
SRCGROUP:prod -> [SRCGROUP:eamv1prod, SRCGROUP:eamv2prod, SRCGROUP:drwv1prod, SRCGROUP:drwv2prod] -> [<eam,drw><01,02>:<v1,v2>:prod]

SRCGROUP:eamv1test -> [eam01:v1:test, eam02:v1:test]
SRCGROUP:eamv1prod -> [eam01:v1:prod, eam02:v1:prod]
SRCGROUP:eamv2test -> [eam01:v2:test, eam02:v2:test]
SRCGROUP:eamv2prod -> [eam01:v2:prod, eam02:v2:prod]
SRCGROUP:drwv1test -> [drw01:v1:test, drw02:v1:test]
SRCGROUP:drwv1prod -> [drw01:v1:prod, drw02:v1:prod]
SRCGROUP:drwv2test -> [drw01:v2:test, drw02:v2:test]
SRCGROUP:drwv2prod -> [drw01:v2:prod, drw02:v2:prod]
```

用户&用户组 与 资源&资源组 之间的关系如下所示。

```text
USER:lvx -> [USERGROUP:it, drw02:v2:prod]
USER:xjw -> [USERGROUP:admin]
USER:ww -> [USERGROUP:quant]

USERGROUP:it -> [drw01:v2:test, drw02:v2:test, SRCGROUP:eamv1test]
USERGROUP:quant -> [SRCGROUP:eamv2prod, SRCGROUP:eamv2test]
USERGROUP:admin -> [*]
USERGROUP:mgr -> [SRCGROUP:v2]
```

The explanation is as follows:

1. The user 'lvx' has access to user group 'it' and resource 'drw02:v2:prod'.
2. The user 'xjw' has access to user group 'admin'.
3. The user 'ww' has access to user group 'quant'.
4. The user group 'it' has access to resource 'drw01:v2:test', resource 'drw02:v2:test' and resource group 'eamv1test'.
5. The user group 'quant' has access to resource group 'eamv2prod'.
6. The user group 'admin' has access to all resources.
7. The resource group 'eam' has access to all resources starting with 'eam', and the resource group 'drw' follows the same principle.
8. The resource group 'prod' has access to all resources ending with 'prod', and the resource group 'test' follows the same principle.
9. The resource group 'v1' has access to all resources marked with 'v1', and the resource group 'v2' follows the same principle.

In summary, the relationships between users and resources are as follows:

- lxv: `drw03:v2:prod`, `eam01:v2:prod`, `<eam,drw><01-03>:<v1,v2>:test`, `<eam,drw><01-03>:v1:<prod,test>`
- xjw: `*`
- ww: `<eam,drw><01-03>:v2:<prod,test>`

Note:
1. To facilitate differentiation between different entities, a prefix will be added to objects indicating their associated entity.
2. `*` represents a wildcard.

