returnValue=`curl -s https://learn.zone01dakar.sn/assets/superhero/all.json`
echo $returnValue  | cat > list.json
family=$(jq -r --arg HERO_ID $HERO_ID '.[]  | select(.id==($HERO_ID|tonumber)) | .connections.relatives' list.json) 
echo -n $family