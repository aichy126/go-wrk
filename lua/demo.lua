request = function()
    local randseed=math.randomseed(os.time())
    local uid = math.random(1, 100000)
    local body = '{"uid":' .. uid .. '}'
    return body,uid
end
