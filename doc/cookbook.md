# soar常用命令
`echo "select physical_cluster_name from blade_tidb" | ./soar -config "./soar.yaml" -log-output=soar.log`

## 输出json
`-report-type json`
echo "select physical_cluster_name from blade_tidb" | ./soar -config "./soar.yaml" -report-type json

## 输出ast  
`-report-type ast`
echo "select physical_cluster_name from blade_tidb" | ./soar -config "./soar.yaml" -report-type ast