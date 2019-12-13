request = function()
    local randseed=math.randomseed(os.time())
    local uid = RandInt(1, 10)
    local body = '{"uid":' .. uid .. '}'
    return body,uid
end
