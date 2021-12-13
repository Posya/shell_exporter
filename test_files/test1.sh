echo '# HELP test1_metric1 First metric'
echo '# TYPE test1_metric1 counter'
echo 'test1_metric1{tag1="t1",tag2="t2"} 1234 1395066363000'

echo '# HELP test1_metric2 First metric'
echo '# TYPE test1_metric2 counter'
echo 'test1_metric2{tag1="t1",tag2="t2"} 5678 1395066363001'

echo 'wrong metric 1234'


